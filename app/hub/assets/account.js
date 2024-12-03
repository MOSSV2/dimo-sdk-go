document.addEventListener("DOMContentLoaded", () => {
  const cardContainer = document.getElementById("cardContainer");
  listAccount()
    .then((data) => {
      data.forEach((meta) => {
        const card = createCard(meta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching accounts:", error));

});

function createCard(meta) {
  const card = document.createElement("div");
  card.className = "card";


  const h2 = document.createElement("h2");
  h2.textContent = meta.Name;
  card.appendChild(h2);

  const p0 = document.createElement("p");
  p0.textContent = `Creation: ${meta.CreatedAt}`;
  card.appendChild(p0);

  const p1 = document.createElement("p");
  p1.textContent = `Needle List: [click to expand]`;
  p1.onclick = () => toNeedle(meta.Name);
  card.appendChild(p1);

  const p2 = document.createElement("p");
  p2.textContent = `Volume List: [click to expand]`;
  p2.onclick = () => toVolume(meta.Name);
  card.appendChild(p2);

  return card;
}

function toNeedle(eaddr) {
  window.location.href = `needle.html?owner=${eaddr}`;
}

function toVolume(eaddr) {
  window.location.href = `volume.html?owner=${eaddr}`;
}


function listAccount() {
  return fetch(`/api/listAccount`)
    .then((response) => response.json())
    .then((data) => {
      return data;
    })
    .catch((error) => console.error("Error fetching accounts:", error))
}

function search() {
  const searchInput = document.getElementById("searchInput").value;
  fetch(`/api/getAccount?owner=${searchInput}`)
    .then((response) => response.json())
    .then((data) => {
      displayResults(data);
    })
    .catch((error) => console.error("Error:", error));
}

function displayResults(data) {
  const resultsElement = document.getElementById("cardContainer");
  resultsElement.innerHTML = ""; // 清空先前的结果

  if (data.length === 0) {
    resultsElement.innerText = "No results found.";
  } else {
    const card = createCard(data);
    cardContainer.appendChild(card);
  }
}