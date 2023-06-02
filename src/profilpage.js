const triangle = document.getElementById('triangle');
const profileImage = document.getElementById('profileImage');
let IdUser = document.getElementById('IdUser');
let IdName = document.getElementById('IdName');
let Message = document.getElementsByClassName('message1');
let Message2 = document.getElementsByClassName('message');
let save = document.getElementsByClassName('save')
let cat = document.getElementsByClassName('cat')
let image = ["morganapdp.png","igypdp.png"]
let imageChanged = 0;
//Message2[0].style.backgroundImage = "url('../asset/bulle.png')"
function Saved(){
    fetch('/savedProfil',{
        method: "POST",
        headers: {"content-type":"application/json"},
        body: JSON.stringify({
            Pictures : imageChanged,
            Name : document.getElementById('IdName').value
        }) 
    })
}
let resp = fetch('/loadUser').then((res) => {
    return res.json()
}).then((d) => {
   // console.log(d[0].Category_id)
    profileImage.src = '../asset/' + image[d.Picture]
    IdUser.innerText = d.Id
    IdName.value = d.Name
});
let resp1 = fetch('/loadPostUser').then((res) => {
    return res.json()
}).then((d) => {
    Message[0].innerText = d[0].Content
    if (d[0].Category_id == 1){
        cat[0].src = '../asset/anime.png'
    }
    if (d[0].Category_id == 2){
        cat[0].src = '../asset/games.png'
    }
    if (d[0].Category_id == 3){
        cat[0].src = '../asset/random.png'
    }
    console.log(d[0])
}).catch(error => {
    console.error(error)
});
triangle.addEventListener('click', () => {
    if (imageChanged+1 == image.length){
        imageChanged = 0
    } else {
        imageChanged++
    }
    profileImage.src = '../asset/' + image[imageChanged]
});
save[0].addEventListener('click', () => {
   Saved();
   location.reload(true);
});