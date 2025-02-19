document.addEventListener("DOMContentLoaded", () => {
  const cardContainer = document.getElementById("cardContainer");

  const urlParams = new URLSearchParams(window.location.search);
  const eaddr = urlParams.get("owner");

  listVolume(eaddr)
    .then((data) => {
      data.forEach((meta) => {
        const card = createCard(meta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching volume:", error));
});

function createCard(meta) {
  const card = document.createElement("div");
  card.className = "card";

  const h2 = document.createElement("h2");
  h2.textContent = meta.Piece;
  card.appendChild(h2);

  const p100 = document.createElement("p");
  p100.textContent = `ChainType: ${meta.ChainType}`;
  card.appendChild(p100);

  const p1 = document.createElement("p");
  p1.textContent = `Owner: ${meta.Owner}`;
  card.appendChild(p1);

  const p2 = document.createElement("p");
  p2.textContent = `Volume: ${meta.File}`;
  card.appendChild(p2)

  const p3 = document.createElement("p");
  p3.textContent = `Creation: ${meta.CreatedAt}`;
  card.appendChild(p3)

  if (meta.TxHash) {
    let rurl = "https://sepolia-optimism.etherscan.io/tx/"
    if (meta.ChainType == "bnb-testnet") {
      rurl = "https://testnet.bscscan.com/tx/"
    } else if (meta.ChainType == "opbnb-testnet") {
      rurl = "https://opbnb-testnet.bscscan.com/tx/"
    }
    const p5 = document.createElement("p");
    p5.innerHTML = "TxHash: <a href='" + rurl + meta.TxHash + "' target='_blank'>" + meta.TxHash + "</a >";
    card.appendChild(p5);
  }

  return card;
}

function listVolume(eaddr) {
  if (eaddr) {
    return fetch(`/api/listVolume?owner=${eaddr}`)
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => console.error("Error fetching volume:", error))
  } else {
    return fetch(`/api/listVolume`)
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => console.error("Error fetching volume:", error))
  }
}


function search() {
  const searchInput = document.getElementById("searchInput").value;
  fetch(`/api/getVolume?owner=${searchInput}`)
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
    data.forEach((meta) => {
      const card = createCard(meta);
      resultsElement.appendChild(card);
    });
  }
}