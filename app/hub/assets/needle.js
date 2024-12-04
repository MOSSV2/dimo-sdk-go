const urlParams = new URLSearchParams(window.location.search);
const eaddr = urlParams.get("owner");

const itemsPerPage = 10;
let currentPage = 1;

function renderPage() {
  const cardContainer = document.getElementById("cardContainer");
  // Clear the current content
  cardContainer.innerHTML = '';

  const startIndex = (currentPage - 1) * itemsPerPage;
  listNeedle(eaddr, startIndex, itemsPerPage)
    .then((data) => {
      data.forEach((fileMeta) => {
        const card = createCard(fileMeta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching needle:", error));

  // Update the page number
  document.getElementById('pageNumber').innerText = `Page: ${currentPage}`;
}

// Handle "Previous" button click
document.getElementById('prevBtn').addEventListener('click', function () {
  if (currentPage > 1) {
    currentPage--;
    renderPage();
  }
});

// Handle "Next" button click
document.getElementById('nextBtn').addEventListener('click', function () {
  currentPage++;
  renderPage();
});

// Initial page render
renderPage();

function createCard(meta) {
  const card = document.createElement("div");
  card.className = "card";
  //card.onclick = () => download(meta.Owner, meta.Name);

  const h2 = document.createElement("h2");
  h2.textContent = meta.Name;
  card.appendChild(h2);

  const p5 = document.createElement("p");
  p5.textContent = `Owner: ${meta.Owner}`;
  card.appendChild(p5);

  const p0 = document.createElement("p");
  p0.textContent = `Volume: ${meta.File}`;
  card.appendChild(p0);

  const p1 = document.createElement("p");
  p1.textContent = `Start: ${meta.Start}`;
  card.appendChild(p1);

  const p2 = document.createElement("p");
  p2.textContent = `Size: ${meta.Size} bytes`;
  card.appendChild(p2);

  const p21 = document.createElement("p");
  p21.textContent = `Creation: ${meta.CreatedAt}`;
  card.appendChild(p21);

  if (meta.TxHash) {
    const p3 = document.createElement("p");
    p3.textContent = `Piece: ${meta.Piece}`;
    card.appendChild(p3);
  }

  if (meta.TxHash) {
    const p5 = document.createElement("p");
    p5.innerHTML = "TxHash: <a href='https://sepolia-optimism.etherscan.io/tx/" + meta.TxHash + "' target='_blank'>" + meta.TxHash + "</a >";
    card.appendChild(p5);
  }

  const p6 = document.createElement("p");
  //p6.textContent = `Content: [click to show]`;
  //p6.onclick = () => download(meta.Owner, meta.Name);
  p6.innerHTML = "Content: <a href='content.html?owner=" + meta.Owner + "&name=" + meta.Name + "' target='_blank'> [click to show] </a >";
  card.appendChild(p6);

  return card;
}

function listNeedle(eaddr, offset, length) {
  if (eaddr) {
    return fetch(`/api/listNeedle?owner=${eaddr}&offset=${offset}&length=${length}`)
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => console.error("Error fetching needles:", error))
  } else {
    return fetch(`/api/listNeedle?offset=${offset}&length=${length}`)
      .then((response) => response.json())
      .then((data) => {
        return data;
      })
      .catch((error) => console.error("Error fetching needles:", error))
  }
}

function search() {
  const searchInput = document.getElementById("searchInput").value;
  fetch(`/api/getNeedle?name=${searchInput}`)
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
      cardContainer.appendChild(card);
    });
  }
}


