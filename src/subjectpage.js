let url = window.location.href
let id_cat = url.substr(url.length -1)
let title = document.querySelector('.title')
let cat = document.querySelector('.subtitles')
let c = 0;
if (id_cat == "1"){
    title.innerHTML = `<img src="../asset/anime.png" alt="Image">
    <p>Animes</p>`
} else if (id_cat == "2"){
    title.innerHTML = `<img src="../asset/games.png" alt="Image">
    <p>Jeux Videos</p>`
} else if (id_cat == "3"){
title.innerHTML = `<img src="../asset/random.png" alt="Image">
<p>Autre</p>`
}
// console.log(id_cat)
window.addEventListener('DOMContentLoaded', function() {
    var categories = document.querySelectorAll('.categories .subtitle');
    
    for (var i = 0; i < categories.length; i++) {
        if (i >= 4) {
            categories[i].style.display = 'none';
        }
    }
});
let resp = fetch('/loadSubjects?id=' + id_cat).then((res) => {
    return res.json()
}).then((d) => {
   console.log(d)
   for (c = 0; c < d.length; c++){
    let sub = document.createElement('div')
    sub.className = 'subtitle'
    sub.innerHTML = `<p>${d[c].Subject}</p>`
    cat.appendChild(sub)
   }
});

document.querySelector('.tobecontinued img.normal').addEventListener('click', function() {
    window.location.href = 'page-suivante.html';
});

document.querySelector('.tobecontinued img.flipped').addEventListener('click', function() {
    window.history.back();
});
