// Variables pour garder une trace de l'état du bouton
var isLiked = false;
var likeCount = 0;

// Fonction pour activer/désactiver le "J'aime"
function toggleLike() {
  // Récupérer l'élément du compteur de likes
  var likeCountElement = document.getElementById('poucebloClics');

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

// Variables pour garder une trace de l'état du bouton
var isDisliked = false;
var dislikeCount = 0;

// Fonction pour activer/désactiver le "J'aime pas"
function toggleDislike() {
  // Récupérer l'élément du compteur de dislikes
  var dislikeCountElement = document.getElementById('poucerougeClics');

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

var isFuck = false;
var fuckCount = 0;

// Fonction pour activer/désactiver le "fuck"
function toggleFuck() {
  // Récupérer l'élément du compteur de fuck
  var fuckCountElement = document.getElementById('fuckClics');

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
    let message = inputElement.value;

    // Créer bulle envoyé avec texte dedans
    let bubbleElement = document.createElement('div');
    bubbleElement.classList.add('conversation');
    bubbleElement.innerHTML = `
        <img class="pdp" src="/asset/igypdp.png">
        <img class="bulletext" src="/asset/bulles.png">
        <span id="poucebloClics" class="pouceblo" style="left: 74vh;font-size: 200%;">0</span>
        <input type="image" id="poucehaut" class="pouceblo" src="/asset/cool.png" onclick="toggleLike()"></input>
        <span id="poucerougeClics" class="poucerouge" style="left: 85vh;font-size: 200%;">0</span>
        <input type="image" id="poucebas" class="poucerouge" src="/asset/ugh.png" onclick="toggleDislike()"></input>
        <span id="fuckClics" class="fuck" style="left: 95vh;font-size: 200%;">0</span>
        <input type="image" id="fucker" class="fuck" src="/asset/duh.png" onclick="toggleFuck()"></input>
    `;

    // Ajouter le message à la bulle de discussion
    bubbleElement.appendChild(document.createTextNode(message));

    // Ajouter la bulle de discussion à l'élément de la conversation
    conversationElement.appendChild(bubbleElement);

    //reinitialisation zone texte
    inputElement.value = '';
  }
});
