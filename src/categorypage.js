let cat = document.getElementsByClassName('subtitle')
let vol = document.getElementById('vol')
let closet = document.getElementById('fishing')
let playing = false;
function Sub(i){
    fetch('/subject2',{
        method: "POST",
        headers: {"content-type":"application/json"},
        body: JSON.stringify({
            Signal : i
        }) 
    })
}
cat[0].addEventListener('click', () =>{
    Sub(1)
})
cat[1].addEventListener('click', () =>{
    Sub(2)
})
cat[2].addEventListener('click', () =>{
    Sub(3)
})
vol.addEventListener('click', () => {
    if (playing){
      closet.pause()
      playing = false;
    } else {
      closet.play()
      playing = true;
    }
  });