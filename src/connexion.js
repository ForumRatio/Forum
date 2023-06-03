let button = document.getElementsByClassName('button')
let input = document.querySelectorAll('input')
function Check(){
    fetch('/checkUser',{
        method: "POST",
        headers: {"content-type":"application/json"},
        body: JSON.stringify({
            Username : input[0].value,
            Password: input[1].value
        }) 
    })
    .then((re) => {
        console.log(re)
        if (re.redirected){
            window.location="/categorypage"
        }
    })
}
button[0].addEventListener('click', () => {
    Check();
    // window.location="/categorypage"
 });
console.log(button[0])