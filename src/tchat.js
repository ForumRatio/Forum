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

let inputElement = document.querySelector('.comecrit');
let conversationElement = document.querySelector('.conversation');
let bubbleIdCounter = 1;

// Fonction pour activer/désactiver le "J'aime"
function toggleLike(element) {
  let bubbleElement = element.closest('.conv');
  let likeCountElement = bubbleElement.querySelector('.poucebloClics');
  let likeCount = parseInt(likeCountElement.getAttribute('data-count'));

  if (element.classList.contains('active')) {
    // Annuler le like
    likeCount--;
    element.classList.remove('active');
  } else {
    // Ajouter le like
    likeCount++;
    element.classList.add('active');
  }

  likeCountElement.setAttribute('data-count', likeCount);
  likeCountElement.textContent = likeCount;
}

// Fonction pour activer/désactiver le "J'aime pas"
function toggleDislike(element) {
  let bubbleElement = element.closest('.conv');
  let dislikeCountElement = bubbleElement.querySelector('.poucerougeClics');
  let dislikeCount = parseInt(dislikeCountElement.getAttribute('data-count'));

  if (element.classList.contains('active')) {
    // Annuler le dislike
    dislikeCount--;
    element.classList.remove('active');
  } else {
    // Ajouter le dislike
    dislikeCount++;
    element.classList.add('active');
  }

  dislikeCountElement.setAttribute('data-count', dislikeCount);
  dislikeCountElement.textContent = dislikeCount;
}

// Fonction pour activer/désactiver le "fuck"
function toggleFuck(element) {
  let bubbleElement = element.closest('.conv');
  let fuckCountElement = bubbleElement.querySelector('.fuckClics');
  let fuckCount = parseInt(fuckCountElement.getAttribute('data-count'));

  if (element.classList.contains('active')) {
    // Annuler le fuck
    fuckCount--;
    element.classList.remove('active');
  } else {
    // Ajouter le fuck
    fuckCount++;
    element.classList.add('active');
  }

  fuckCountElement.setAttribute('data-count', fuckCount);
  fuckCountElement.textContent = fuckCount;
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
let bubbleId = `bubble-${bubbleIdCounter}`;
bubbleElement.id = bubbleId;
bubbleElement.classList.add('conv');
bubbleElement.innerHTML = `
<img class="pdp" src="/asset/igypdp1.png">
        <img class="bulletext" src="/asset/bulles.png">
        <div class="container">
          <div class="blo">
          <span id="poucebloClics-${bubbleIdCounter}" class="poucebloClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="pouceblo" src="/asset/cool.png" onclick="toggleLike(this)"></input>
          </div>
          <div class="rouge">
          <span id="poucerougeClics-${bubbleIdCounter}" class="poucerougeClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="poucerouge" src="/asset/ugh.png" onclick="toggleDislike(this)"></input>
          </div>
          <div class="fucks">
          <span id="fuckClics-${bubbleIdCounter}" class="fuckClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="fuck" src="/asset/duh.png" onclick="toggleFuck(this)"></input>
          </div>
        </div>
        <input class="nomutilisateur" type="text" placeholder="Nom">
`;
bubbleIdCounter++;
// Ajouter le message à la bulle de discussion
bubbleElement.appendChild(message);

// Ajouter la bulle de discussion à l'élément de la conversation
conversationElement.appendChild(bubbleElement);

//reinitialisation zone texte
inputElement.value = '';
  }
  console.log(document.querySelectorAll('.pouceblo'))
}
