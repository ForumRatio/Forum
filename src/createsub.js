let elt = document.querySelectorAll('input');
elt[0].maxLength = 30;
elt[1].maxLength = 30;
let button = document.querySelector('.save')
let select = document.getElementById("category_select")
    select.addEventListener('change',function () {
        select.value = this.value;
        console.log(select.value)
})
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