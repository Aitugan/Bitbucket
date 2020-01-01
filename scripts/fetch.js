fetch('../scripts/product.json')
    .then(function(response) {
        return response.json();
    })
    .then(function(data) {
        appendData(data);
    })
    .catch(function(err) {
        console.log(err);
    });

function appendData(data) {
    let dataTitle = document.getElementById("dataTitle");
    var h2 = document.createElement("h2");
    h2.innerHTML = data.title;
    dataTitle.appendChild(h2);
    document.getElementById("dataP1").innerHTML = data.dataP1;
    document.getElementById("dataP2").innerHTML = data.dataP2;
    document.getElementById("hours").innerHTML = data.hours + "hours video on demand";
    document.getElementById("hours").appendChild(h);
}