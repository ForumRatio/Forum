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