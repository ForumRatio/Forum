let url = window.location.href
let id_cat = url.substr(url.length -1)
let title = document.querySelector('.title')
let cat = document.querySelector('.subtitles')
let c = 0;
let e = 0;
let start = 1;
if (id_cat == "1"){
    title.innerHTML = `<img src="../asset/anime.png" alt="Image">
    <p>Animes</p>
    <img src="../asset/vol.png" id="vol">`
} else if (id_cat == "2"){
    title.innerHTML = `<img src="../asset/games.png" alt="Image">
    <p>Jeux Videos</p>
    <img src="../asset/vol.png" id="vol">`
} else if (id_cat == "3"){
title.innerHTML = `<img src="../asset/random.png" alt="Image">
    <p>Autre</p>
    <img src="../asset/vol.png" id="vol">`
}
// let vol = document.querySelector('#vol');
let playing = false;
let closet = document.querySelector('#fishing');
window.addEventListener('DOMContentLoaded', function() {
    resp()
    document.querySelector('.tobecontinued img.normal').addEventListener('click', function() {
        // window.location.href = 'page-suivante.html';
        nextPage()
    });
    
    document.querySelector('.tobecontinued img.flipped').addEventListener('click', function() {
        // window.history.back();
        previousPage()
    });
    document.querySelector('#vol').addEventListener('click', function() {
        if (playing) {
            closet.pause();
            playing = false;
          } else {
            closet.play();
            playing = true;
          }
    });
});

function resp(){
    fetch('/loadSubjects?id=' + id_cat).then((res) => {
        return res.json()
    }).then((d) => {
       console.log(d)
       e = start * 3
       cat.innerHTML = ``
       if (start * 3 < d.length){
        for (c = (start - 1) *3; c < start *3; c++){
            let sub = document.createElement('div')
            sub.className = 'subtitle'
            sub.innerHTML = `<a href="/tchat?id=${d[c].Id}&cat=${d[c].Category_id}"><p class="subName">${d[c].Subject}</p></a>`
            cat.appendChild(sub)
           }
       } else {
        for (c = (start - 1) *3; c < d.length; c++){
            let sub = document.createElement('div')
            sub.className = 'subtitle'
            sub.innerHTML = `<a href="/tchat?id=${d[c].Id}&cat=${d[c].category_id}"><p class="subName">${d[c].Subject}</p></a>`
            cat.appendChild(sub)
           }
       }
    });
} 
function nextPage(){
    if (start < 10){
        start++;
        resp();
    } 
}
function previousPage(){
    if (start > 1){
        start--
        resp()
    }
}

console.log()