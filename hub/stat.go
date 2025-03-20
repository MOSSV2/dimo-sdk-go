package hub

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) addStat(g *gin.RouterGroup) {
	g.Group("/").POST("/stat", s.getStatByPost)
	g.Group("/").GET("/stat", s.getStatByGet)
}

// @Summary Get statistics by POST
// @Description Get daily statistics for accounts, needles, and volumes using POST method
// @Tags statistics
// @Accept x-www-form-urlencoded
// @Produce json
// @Param count formData int false "Number of days to return (default: 30)"
// @Success 200 {array} types.Stat "Array of daily statistics"
// @Failure 400 {object} map[string]string "Invalid count parameter"
// @Router /api/stat [post]
func (s *Server) getStatByPost(c *gin.Context) {
	count := 30
	countStr := c.PostForm("count")
	countInt, err := strconv.Atoi(countStr)
	if err == nil && countInt > 0 {
		count = countInt
	}
	stats := s.statManager.GetStats(count)
	c.JSON(200, stats)
}

// @Summary Get statistics by GET
// @Description Get daily statistics for accounts, needles, and volumes using GET method
// @Tags statistics
// @Accept json
// @Produce json
// @Param count query int false "Number of days to return (default: 30)"
// @Success 200 {array} types.Stat "Array of daily statistics"
// @Failure 400 {object} map[string]string "Invalid count parameter"
// @Router /api/stat [get]
func (s *Server) getStatByGet(c *gin.Context) {
	count := 30
	countStr := c.Query("count")
	countInt, err := strconv.Atoi(countStr)
	if err == nil && countInt > 0 {
		count = countInt
	}
	stats := s.statManager.GetStats(count)
	c.JSON(200, stats)
}

// StatManager manages daily statistics for accounts, needles, and volumes
type StatManager struct {
	mu       sync.RWMutex
	db       *gorm.DB
	stats    map[string]*types.Stat // key: YYYY-MM-DD
	lastDay  time.Time
	stopChan chan struct{}
}

// NewStatManager creates a new StatManager instance
func NewStatManager(db *gorm.DB) *StatManager {
	return &StatManager{
		db:       db,
		stats:    make(map[string]*types.Stat),
		stopChan: make(chan struct{}),
	}
}

// Start initializes the StatManager with historical data and starts the background update routine
func (sm *StatManager) Start(ctx context.Context) {
	// Initialize data for the last 30 days
	now := time.Now().UTC()
	// Align to day boundary
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	startDay := now.AddDate(0, 0, -31)
	sm.updateStat(startDay)

	for i := 30; i >= 0; i-- {
		day := now.AddDate(0, 0, -i)
		sm.updateDailyStat(day)
	}
	sm.lastDay = now

	// Start background update routine
	go sm.run(ctx)
}

// Stop stops the background update routine
func (sm *StatManager) Stop() {
	close(sm.stopChan)
}

// run executes the background update routine
func (sm *StatManager) run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-sm.stopChan:
			return
		case <-ticker.C:
			now := time.Now().UTC()
			// Align to day boundary
			now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

			// Update previous day's data if date changed
			if now.Day() != sm.lastDay.Day() {
				sm.updateDailyStat(sm.lastDay)
				sm.lastDay = now
			}
			// Update current day's data
			sm.updateDailyStat(now)
		}
	}
}

// updateStat updates statistics for before a day
func (sm *StatManager) updateStat(t time.Time) {
	// Ensure time is aligned to day boundary
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	day := t.Format("2006-01-02")
	nextDay := t.Add(24 * time.Hour)

	// Get total counts up to the specified day
	var totalAccounts []types.Account
	sm.db.Model(&types.Account{}).Where("created_at < ?", nextDay).Find(&totalAccounts)
	totalAccountsCount := int64(len(totalAccounts))

	var totalNeedles []types.Needle
	sm.db.Model(&types.Needle{}).Where("created_at < ?", nextDay).Find(&totalNeedles)
	totalNeedlesCount := int64(len(totalNeedles))

	var totalVolumes []types.Volume
	sm.db.Model(&types.Volume{}).Where("created_at < ?", nextDay).Find(&totalVolumes)
	totalVolumesCount := int64(len(totalVolumes))

	stat := &types.Stat{
		Day:           t,
		TotalAccounts: totalAccountsCount,
		TotalNeedles:  totalNeedlesCount,
		TotalVolumes:  totalVolumesCount,
	}

	sm.mu.Lock()
	sm.stats[day] = stat
	sm.mu.Unlock()
	logger.Info("stat: ", stat)
}

// updateDayStat updates statistics for a specific day
func (sm *StatManager) updateDailyStat(t time.Time) {
	// Ensure time is aligned to day boundary
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	day := t.Format("2006-01-02")
	nextDay := t.Add(24 * time.Hour)

	// Get total counts up to the specified day
	var dailyAccounts []types.Account
	sm.db.Model(&types.Account{}).Where("created_at < ? and created_at >= ?", nextDay, t).Find(&dailyAccounts)
	dailyAccountsCount := int64(len(dailyAccounts))

	var dailyNeedles []types.Needle
	sm.db.Model(&types.Needle{}).Where("created_at < ? and created_at >= ?", nextDay, t).Find(&dailyNeedles)
	dailyNeedlesCount := int64(len(dailyNeedles))

	var dailyVolumes []types.Volume
	sm.db.Model(&types.Volume{}).Where("created_at < ? and created_at >= ?", nextDay, t).Find(&dailyVolumes)
	dailyVolumesCount := int64(len(dailyVolumes))

	sm.mu.Lock()
	stat, ok := sm.stats[day]
	if !ok {
		stat = &types.Stat{
			Day: t,
		}
		sm.stats[day] = stat
	}
	stat.DailyAccounts = dailyAccountsCount
	stat.DailyNeedles = dailyNeedlesCount
	stat.DailyVolumes = dailyVolumesCount

	prevDayTime := t.AddDate(0, 0, -1)
	prevDay := prevDayTime.Format("2006-01-02")
	prevDayStat, ok := sm.stats[prevDay]
	if !ok {
		prevDayStat = &types.Stat{
			Day: prevDayTime,
		}
		sm.stats[prevDay] = prevDayStat
	}
	stat.TotalAccounts = prevDayStat.TotalAccounts + dailyAccountsCount
	stat.TotalNeedles = prevDayStat.TotalNeedles + dailyNeedlesCount
	stat.TotalVolumes = prevDayStat.TotalVolumes + dailyVolumesCount

	sm.mu.Unlock()
	logger.Info("stat: ", stat)
}

// GetStats returns a copy of all statistics
func (sm *StatManager) GetStats(count int) []types.Stat {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// Return a copy to prevent external modification
	// sort the result by day desc
	if count > len(sm.stats) {
		count = len(sm.stats)
	}
	result := make([]types.Stat, 0, count)
	now := time.Now().UTC()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	for i := 0; i < count; i++ {
		day := now.AddDate(0, 0, -i)
		stat, ok := sm.stats[day.Format("2006-01-02")]
		if !ok {
			continue
		}
		result = append(result, *stat)
	}

	return result
}
