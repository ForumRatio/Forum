// variables pour garder une trace de l'état du bouton
let isLiked = false;
let likeCount = 0;

// Fonction pour activer/désactiver le "J'aime"
function toggleLike() {
  // Récupérer l'élément du compteur de likes
  let likeCountElement = document.getElementById('poucebloClics');

  // Mettre à jour l'état du bouton et le compteur de likes
  if (isLiked) {
    likeCount--;
  } else { 
    likeCount++;
  }
  
  // Mettre à jour l'affichage du compteur de likes
  likeCountElement.innerHTML = likeCount.toString();

  // Inverser l'état du bouton
  isLiked = !isLiked;
}

// variables pour garder une trace de l'état du bouton
let isDisliked = false;
let dislikeCount = 0;

// Fonction pour activer/désactiver le "J'aime pas"
function toggleDislike() {
  // Récupérer l'élément du compteur de dislikes
  let dislikeCountElement = document.getElementById('poucerougeClics');

  // Mettre à jour l'état du bouton et le compteur de dislikes
  if (isDisliked) {
    dislikeCount--;
  } else {
    dislikeCount++;
  }

  // Mettre à jour l'affichage du compteur de dislikes
  dislikeCountElement.innerHTML = dislikeCount.toString();

  // Inverser l'état du bouton
  isDisliked = !isDisliked;
}

let isFuck = false;
let fuckCount = 0;

// Fonction pour activer/désactiver le "fuck"
function toggleFuck() {
  // Récupérer l'élément du compteur de fuck
  let fuckCountElement = document.getElementById('fuckClics');

  // Mettre à jour l'état du bouton et le compteur de fuck
  if (isFuck) {
    fuckCount--;
  } else {
    fuckCount++;
  }

  // Mettre à jour l'affichage du compteur de fuck
  fuckCountElement.innerHTML = fuckCount.toString();

  // Inverser l'état du bouton
  isFuck = !isFuck;
}



let inputElement = document.querySelector('.comecrit');
let conversationElement = document.querySelector('.conversation');

// clic
inputElement.addEventListener('keyup', function(event) {
  // si espace alors envoie
  if (event.keyCode === 13) {
    // Récupérer le contenu du champ de saisie
    let message = document.createElement('p');
    message.className=('message')
    message.innerHTML = inputElement.value
    // Créer bulle envoyé avec texte dedans
    let bubbleElement = document.createElement('div');
    bubbleElement.classList.add('conv');
    bubbleElement.innerHTML = `
        <img class="pdp" src="/asset/igypdp.png">
        <img class="bulletext" src="/asset/bulles.png">
        <div class="container">
          <div class="blo">
            <span id="poucebloClics" class="pouceblo" style="font-size: 200%;">0</span>
            <input type="image" id="poucehaut" class="pouceblo" src="/asset/cool.png" onclick="toggleLike()"></input>
          </div>
          <div class="rouge">
            <span id="poucerougeClics" class="poucerouge" style="font-size: 200%;">0</span>
            <input type="image" id="poucebas" class="poucerouge" src="/asset/ugh.png" onclick="toggleDislike()"></input>
          </div>
          <div class="fucks">
            <span id="fuckClics" class="fuck" style="font-size: 200%;">0</span>
            <input type="image" id="fucker" class="fuck" src="/asset/duh.png" onclick="toggleFuck()"></input>
          </div>
        </div>
    `;

    // Ajouter le message à la bulle de discussion
    bubbleElement.appendChild(message);

    // Ajouter la bulle de discussion à l'élément de la conversation
    conversationElement.appendChild(bubbleElement);

    //reinitialisation zone texte
    inputElement.value = '';
  }
});
