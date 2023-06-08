const triangle = document.getElementById('triangle');
const profileImage = document.getElementById('profileImage');
let IdUser = document.getElementById('IdUser');
let IdName = document.getElementById('IdName');
let Message = document.getElementsByClassName('message1');
let Message2 = document.getElementsByClassName('message');
let save = document.getElementsByClassName('save')
let cat = document.getElementsByClassName('cat')
let table = document.querySelector('tbody')
let vol = document.getElementById('vol')
let sunset = document.getElementById('sunset')
let closet = document.getElementById('closet')
let image = ["morganapdp.png","igypdp.png"]
let playing = false;
let delete1;
let message1;
let edit1;
let imageChanged = 0;
let c = 0;
let d = 0;

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
    // console.log(d)
    let im;
    for (c = 0; c < d.length; c++){
        if (d[c].Category_id == 1){
                im = '../asset/anime.png'
            }
        if (d[c].Category_id == 2){
                im = '../asset/games.png'
            }
        if (d[c].Category_id == 3){
                im = '../asset/random.png'
            }
        let cat1 = document.createElement('tr')
        cat1.innerHTML = `<td>
        <div class="category">
          <img src="${im}" alt="Category 1" class="cat">
        </div>
      </td>
      <td class="message">
        <p class="message1">${d[c].Content}</p>
      </td>
      <td>
        <div class="category">
          <img src="../asset/settings.png" alt="Category 3" a href="#" class="edit">
        </div>
      </td>
      <td>
        <div class="category">
          <img src="../asset/delete.png" alt="Category 4" a href="#" class="delete">
        </div>
      </td>`
      table.appendChild(cat1)
    }
    delete1 = document.querySelectorAll('.delete')
    message1 = document.querySelectorAll('.message1')
    edit1 = document.querySelectorAll('.edit')
    for (d = 0; d < message1.length; d++){
      let e = d
      let message2 = message1[d]
      delete1[d].addEventListener('click', ()=>{
        console.log(message2)
        fetch('/deletePost',{
          method: "POST",
          headers: {"content-type":"application/json"},
          body: JSON.stringify({
              Check : message2.innerHTML
          }) 
      }).then((res) => {
        if (res.redirected){
          // location.reload(true);
          let cat2 = document.querySelectorAll('tr')
          table.removeChild(cat2[e])
        }
      })
      })
      edit1[d].addEventListener('click', ()=>{
        let text;
        let content = prompt("Please enter your new sentence:", message2.innerHTML);
        if (content == null || content == "") {
          // text = "User cancelled the prompt.";
        } else {
          text = content;
          fetch('/editPost',{
            method: "POST",
            headers: {"content-type":"application/json"},
            body: JSON.stringify({
                Check : message2.innerHTML,
                NS : text
            }) 
        })
        .then((res) => {
          if (res.redirected){
            message2.innerHTML = text
          }
        })
      }
    })
  }
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
  //  location.reload(true);
});
vol.addEventListener('click', () => {
  // console.log(closet)
  if (playing){
    closet.pause()
    playing = false;
  } else {
    closet.play()
    playing = true;
  }
});