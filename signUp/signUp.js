console.log("JS loaded")

const port = ":8081"
// const url = "/signup"

var inputForm = document.getElementById("myForm")

inputForm.addEventListener("submit",(e) => {
    e.preventDefault

    const formdata = new FormData(inputForm)

    fetch(port + url, {method:"POST",body:formdata
}).then(
    response => response.text() 
    ).then(
        (data) => {
            console.log(data)
            document.getElementById("errors").innerHTML =data
        }
    ).catch(
        error => console.error(error)
    )
}
)