document.addEventListener("DOMContentLoaded", () => {
  const cardContainer = document.getElementById("cardContainer");

  const urlParams = new URLSearchParams(window.location.search);
  const eaddr = urlParams.get("owner");
  const name = urlParams.get("name");

  download(eaddr, name)
    .then((data) => {
      data.forEach((fileMeta) => {
        const card = createCard(fileMeta);
        cardContainer.appendChild(card);
      });
    })
    .catch((error) => console.error("Error fetching fileMeta:", error));
});


function download(owner, name) {
  fetch(`/api/download?owner=${owner}&name=${name}`)
    .then((response) => response.text())
    .then((data) => {
      const card = createDataCard(owner, name, data);
      cardContainer.appendChild(card);
    })
    .catch((error) => console.error("Error:", error));
}

function createDataCard(owner, name, meta) {
  const card = document.createElement("div");
  card.className = "card";

  const h2 = document.createElement("h2");
  h2.textContent = name;
  card.appendChild(h2);

  const p5 = document.createElement("p");
  p5.textContent = `Owner: ${owner}`;
  card.appendChild(p5);

  const p0 = document.createElement("p");
  p0.innerHTML = "<p>Content:</p>";
  p0.innerHTML += "<pre>" + formatString(meta) + "</pre>"
  //p0.innerHTML += "</pre >"
  card.appendChild(p0);

  return card
}

function formatString(str) {
  try {
    o = JSON.parse(str);
    return JSON.stringify(o, null, 2);
  } catch (e) {
    return str;
  }
}