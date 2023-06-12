let elt = document.querySelectorAll('input');
elt[0].maxLength = 30;
elt[1].maxLength = 50;
let button = document.querySelector('.save')
let select2 = document.getElementById("category_select")
    select2.addEventListener('change',function () {
        select2.value = this.value;
        console.log(select2.value)
})
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

function Saved1(){
    console.log(elt[0].value,select.value)
    fetch('/savedSub',{
        method: "POST",
        headers: {"content-type":"application/json"},
        body: JSON.stringify({
            Subject : elt[0].value,
            Question : elt[1].value,
            Category_id : parseInt(select.value,10)
        }) 
    }).then((res) => {
        if (res.redirected){
            window.location='/categorypage'
        }
    })
}
button.addEventListener('click', () => {
    Saved1()
})