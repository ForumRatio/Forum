let button = document.getElementsByClassName('button')
let input = document.querySelectorAll('input')
function Create(){
    if (input[2].value === input[3].value) {
        fetch('/createUser',{
            method: "POST",
            headers: {"content-type":"application/json"},
            body: JSON.stringify({
            Name : input[0].value,
            Cellphone: input[1].value,
            Email: input[4].value,
            Password : input[2].value
            }) 
        })
        .then((re) => {
            console.log(re)
            if (re.redirected){
                window.location="/categorypage"
            }
        })
    }
}
button[0].addEventListener('click', () => {
    console.log(button[0])
    Create();
    // window.location="/categorypage"
 });
