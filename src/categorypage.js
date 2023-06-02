let cat = document.getElementsByClassName('subtitle')
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