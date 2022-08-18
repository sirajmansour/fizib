console.log("JS Loaded")

const url = ":8081"

var inputForm = document.getElementById("inputForm")

inputForm.addEventListener("submit", (e) => {

    //prevent auto submission
    e.preventDefault()

    const formdata = new FormData(inputForm)

    fetch(url, {
        method: "POST", body: formdata,
    }).then(
        
        response => response.text()
    ).then(
        (data) => {
            console.log(data);
            document.getElementById("serverMessageBox").innerHTML = data;
        }
    ).catch(
        error => console.error(error)
    )
})
