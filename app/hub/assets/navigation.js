document.addEventListener('DOMContentLoaded', function () {
  // Fetch the navigation bar content from navigation.html
  fetch('navigation.html')
    .then(response => response.text())
    .then(content => {
      // Get the container element and insert the navigation bar content
      const container = document.getElementById('navigation-container');
      container.innerHTML = content;
    })
    .catch(error => console.error('Error fetching navigation bar:', error));
});