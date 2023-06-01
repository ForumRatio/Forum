const triangle = document.getElementById('triangle');
const profileImage = document.getElementById('profileImage');
let IdUser = document.getElementById('IdUser');
let IdName = document.getElementById('IdName');
let save = document.getElementsByClassName('save')
let imageChanged = 0;
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
    console.log(d.Picture)
    profileImage.src = '../asset/' + image[d.Picture]
    IdUser.innerText = d.Id
    IdName.value = d.Name
    
});
let image = ["morganapdp.png","igypdp.png"]
//profileImage.src = '../asset/' + image[imageChanged]
//console.log(save[0])
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
});