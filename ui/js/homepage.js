const testButton = document.getElementById("test-button")
const buttonTime = document.getElementById("button-time")
const from = document.getElementById("from")
const to = document.getElementById("to")
const askButton = document.getElementById("ask-button")
const askTime = document.getElementById("ask-time")

const apiName = "/v1/get_distance"

askButton.addEventListener("click", function () {
    let data = {
        From: from.value,
        To: to.value,
        Time: new Date().toLocaleString("en-IE"),
    };
    fetch(apiName, {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            console.log(result)
            askTime.textContent = "Measured shorter distance " + result["Duration"] + " departure time " + result["Time"]
        });
    }).catch((error) => {
        console.log(error)
    });
})

askButton.addEventListener("click", function () {
    askTime.textContent = "new button clicked"
})