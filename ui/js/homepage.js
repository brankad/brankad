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
            if (result["RespCode"] == 1){
                askTime.textContent = "Measured shorter distance " + result["Duration"] + " next departure time " + result["Time"]
            } else {
                askTime.textContent = result["Text"]
            }
        });
    }).catch((error) => {
        console.log(error)
    });
})