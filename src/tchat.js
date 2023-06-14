let url = window.location.href.split("&")
let id_sub = url[0].substr(url[0].length -1)
let id_cat = url[1].substr(url[1].length -1)
let icon = document.querySelector('.games')
let title = document.querySelector('.sujet')
let title2 = document.querySelector('.sujet2')
let e = 1;
let z = 0;
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
function toggleLike(element,Pid,Uid) {
  let bubbleElement = document.querySelectorAll('.conv');
  let likeCountElement = bubbleElement[element].querySelector('.poucebloClics');
  fetch('/savedLike',{
    method: "POST",
    headers: {"content-type":"application/json"},
    body: JSON.stringify({
        Post_id : Pid,
        User_id : Uid
    }) 
}).then(e=>{
  fetch('/loadLike?id=' + Pid).then((res3) => {
    return res3.json()
  }).then((i)=>{
    // console.log(i.length)
    likeCountElement.textContent = i.length;
  })
})
}
// Fonction pour activer/désactiver le "J'aime pas"
function toggleDislike(element,Pid,Uid) {
  let bubbleElement = document.querySelectorAll('.conv');
  let likeCountElement = bubbleElement[element].querySelector('.poucerougeClics');
  fetch('/savedDislike',{
    method: "POST",
    headers: {"content-type":"application/json"},
    body: JSON.stringify({
        Post_id : Pid,
        User_id : Uid
    }) 
}).then(e=>{
  fetch('/loadDislike?id=' + Pid).then((res3) => {
    return res3.json()
  }).then((i)=>{
    // console.log(i.length)
    likeCountElement.textContent = i.length;
  })
})
}

// Fonction pour activer/désactiver le "fuck"
function toggleFuck(element,Pid,Uid) {
  let bubbleElement = document.querySelectorAll('.conv');
  let likeCountElement = bubbleElement[element].querySelector('.fuckClics');
  fetch('/savedFuck',{
    method: "POST",
    headers: {"content-type":"application/json"},
    body: JSON.stringify({
        Post_id : Pid,
        User_id : Uid
    }) 
}).then(e=>{
  fetch('/loadFuck?id=' + Pid).then((res3) => {
    return res3.json()
  }).then((i)=>{
    // console.log(i.length)
    likeCountElement.textContent = i.length;
  })
})
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
  fetch('/loadAllUser').then((res2) => {
    return res2.json()
  }).then((r)=>{
    console.log(d)
    console.log(r)
    for (e = 1; e < d.length; e++){
      let user = ""
      let picture = 0;
      for (z = 0; z < r.length; z++){
        if (d[e].User_id == r[z].Id){
          user = r[z].Name
          picture = r[z].Picture
          break;
        }
      }
      z = 0
      let picture2 = ``;
      if(picture == 0){
        picture2 = `<img class="pdp" src="/asset/morganapdp1.png">`
      } else if(picture == 1){
        picture2 = `<img class="pdp" src="/asset/igypdp1.png">`
      }
      sendMessage2(d[e].Content,user,picture2)
    }
// console.log(document.querySelectorAll('.conv'))
    let pouceblo = document.querySelectorAll('.pouceblo')
    let poucebloClics = document.querySelectorAll('.poucebloClics');
    let poucerouge = document.querySelectorAll('.poucerouge')
    let poucerougeClics = document.querySelectorAll('.poucerougeClics');
    let fuck = document.querySelectorAll('.fuck')
    let fuckClics = document.querySelectorAll('.fuckClics');
    for (z = 0; z < pouceblo.length; z++){
      let c = z
      let Pid = d[z+1].Id
      
  fetch('/loadLike?id=' + Pid).then((res2) => {
        return res2.json()
      }).then((w)=>{
      poucebloClics[c].textContent = w.length 
  });
  fetch('/loadDislike?id=' + Pid).then((res2) => {
    return res2.json()
  }).then((w)=>{
  poucerougeClics[c].textContent = w.length 
  });
  fetch('/loadFuck?id=' + Pid).then((res2) => {
    return res2.json()
  }).then((w)=>{
  fuckClics[c].textContent = w.length 
});
  pouceblo[z].addEventListener('click', ()=>{
    fetch('/loadUser').then((res) => {
      return res.json()
  }).then((d) => {
    toggleLike(c, Pid, d.Id)
  });
  })
  poucerouge[z].addEventListener('click', ()=>{
    fetch('/loadUser').then((res) => {
      return res.json()
  }).then((d) => {
    toggleDislike(c, Pid, d.Id)
  });
  })
  fuck[z].addEventListener('click', ()=>{
    fetch('/loadUser').then((res) => {
      return res.json()
  }).then((d) => {
    toggleFuck(c, Pid, d.Id)
  });
  })
}

});
// clic
inputElement.addEventListener('keyup', function(event) {
  // si entrée alors envoie
  if (event.keyCode === 13) {
  fetch('/loadUser').then((res) => {
      return res.json()
  }).then((d) => {
    let Picture3 = ``;
      if(d.Picture == 0){
        Picture3 = `<img class="pdp" src="/asset/morganapdp1.png">`
      } else if(d.Picture == 1){
        Picture3 = `<img class="pdp" src="/asset/igypdp1.png">`
      }
    sendMessage(inputElement.value,d.Name,Picture3);
  });
  }
})
});

function sendMessage(Content,User,Picture){
sendMessage2(Content,User,Picture)
console.log(document.querySelectorAll('.conv'))
SavedPost(Content, id_sub, id_cat)
  }

function sendMessage2(Content,User,Picture){
  // Récupérer le contenu du champ de saisie
  let message = document.createElement('p');
  message.className=('message')
  if (Content >= 34){
    let first = Content.substring(0, Content.length /2);
    let last = Content.substr(Content.length /2)
    message.innerHTML = `${first} <br> ${last}`
  } else {
    message.innerHTML = Content
  }
  if (message.innerHTML != ""){
// Créer bulle envoyé avec texte dedans
let bubbleElement = document.createElement('div');
let bubbleId = `bubble-${bubbleIdCounter}`;
bubbleElement.id = bubbleId;
bubbleElement.classList.add('conv');
bubbleElement.innerHTML = `
${Picture}
        <img class="bulletext" src="/asset/bulles.png">
        <div class="container">
          <div class="blo">
          <span id="poucebloClics-${bubbleIdCounter}" class="poucebloClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="pouceblo" src="/asset/cool.png"></input>
          </div>
          <div class="rouge">
          <span id="poucerougeClics-${bubbleIdCounter}" class="poucerougeClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="poucerouge" src="/asset/ugh.png"></input>
          </div>
          <div class="fucks">
          <span id="fuckClics-${bubbleIdCounter}" class="fuckClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="fuck" src="/asset/duh.png"></input>
          </div>
        </div>
        <input class="nomutilisateur" type="text" placeholder="Nom" value="${User}" disabled>
`;
bubbleIdCounter++;
// Ajouter le message à la bulle de discussion
bubbleElement.appendChild(message);

// Ajouter la bulle de discussion à l'élément de la conversation
conversationElement.appendChild(bubbleElement);

//reinitialisation zone texte
inputElement.value = '';
  }
}

function SavedPost(a,b,c){
  fetch('/savedPost',{
      method: "POST",
      headers: {"content-type":"application/json"},
      body: JSON.stringify({
          Check : a,
          Subject_id : parseInt(b,10),
          Category_id : parseInt(c,10)
      }) 
  })
}