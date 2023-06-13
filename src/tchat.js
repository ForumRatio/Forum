let url = window.location.href.split("&")
let id_sub = url[0].substr(url[0].length -1)
let id_cat = url[1].substr(url[1].length -1)
let icon = document.querySelector('.games')
let title = document.querySelector('.sujet')
let title2 = document.querySelector('.sujet2')
console.log(id_sub,id_cat)
let vol = document.getElementById('vol');
  let closet = document.getElementById('fishing');
  let playing = false;
  let select = document.getElementById('Element');
  // sortSelect(select);
  select.addEventListener('change', function() {
    console.log(select.options)
    if (!playing){
      closet.setAttribute('src',`./asset/music/${this.value}`)
    } else {
      closet.pause();
      closet.setAttribute('src',`./asset/music/${this.value}`)
      closet.load();
      closet.play();
    } 
  });

  vol.addEventListener('click', function() {
    if (playing) {
      closet.pause();
      playing = false;
    } else {
      closet.play();
      playing = true;
    }
  });
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
let resp2 = fetch('/loadSubjects2?id=' + id_sub).then((res) => {
  return res.json()
}).then((d) => {
  if (d.Category_id == 1){
    icon.src="../asset/anime.png"
  } else if (d.Category_id == 2){
    icon.src="../asset/games.png"
  } else if (d.Category_id == 3){
    icon.src="../asset/random.png"
  }
title.value = d.Subject
});

let resp = fetch('/loadPost?id=' + id_sub).then((res) => {
  return res.json()
}).then((d) => {
  if (d[0].Content.length >= 30){
    let first = d[0].Content.substring(0, d[0].Content.length /2);
    let last = d[0].Content.substr(d[0].Content.length /2)
    title2.value = `${first} ${last}`
  } else {
    title2.value = d[0].Content
  }
  console.log(d)
});

let inputElement = document.querySelector('.comecrit');
let conversationElement = document.querySelector('.conversation');

// clic
inputElement.addEventListener('keyup', function(event) {
  // si espace alors envoie
  if (event.keyCode === 13) {
  sendMessage();
  }
});

function sendMessage(){
  // Récupérer le contenu du champ de saisie
  let message = document.createElement('p');
  message.className=('message')
  if (inputElement.value.length >= 34){
    let first = inputElement.value.substring(0, inputElement.value.length /2);
    let last = inputElement.value.substr(inputElement.value.length /2)
    message.innerHTML = `${first} <br> ${last}`
  } else {
    message.innerHTML = inputElement.value
  }
  if (message.innerHTML != ""){
// Créer bulle envoyé avec texte dedans
let bubbleElement = document.createElement('div');
bubbleElement.classList.add('conv');
bubbleElement.innerHTML = `
<img class="pdp" src="/asset/igypdp1.png">
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
        <input class="nomutilisateur" type="text" placeholder="Nom">
`;

// Ajouter le message à la bulle de discussion
bubbleElement.appendChild(message);

// Ajouter la bulle de discussion à l'élément de la conversation
conversationElement.appendChild(bubbleElement);

//reinitialisation zone texte
inputElement.value = '';
  }
  console.log(document.querySelectorAll('.pouceblo'))
}
