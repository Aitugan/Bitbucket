let pass = document.getElementById("pass");
pass.addEventListener('keyup', function() {
    strength();
})

function strength() {
    let val = document.getElementById("pass").value;
    let status = document.getElementById("length");
    let label = document.getElementById("head");
    let containsIntegers = false;
    // let containsSymbols = false;


    // let prReg1 = new RegExp("!@#$%^&*()_+-");
    // if (prReg1.test(val)) {
    //     containsSymbols = true;
    // }
    if (/\d/.test(val)) {
        containsIntegers = true;
    }

    // console.log(containsIntegers);
    // console.log(containsSymbols)
    status.style.display = "none";
    label.style.display = "none";
    if (val != "") {
        status.style.display = "unset";
        label.style.display = "unset";

        if (val.length <= 5) {
            status.innerHTML = "WEAK";
            status.style.padding = "5px 15px";
            status.style.background = "linear-gradient(to left, rgb(128, 0, 4) 0%, #ae00ff 100%)";
            status.style.color = "white";
        } else
        if (val.length > 5 && val.length <= 10) {
            status.innerHTML = "GOOD";
            status.style.padding = "5px 25px";
            status.style.background = "linear-gradient(to left, rgb(0, 21, 88) 0%, #00b6b6 100%)";
            status.style.color = "white";
        } else
        if (val.length > 10 && containsIntegers) { //&& containsSymbols) {
            status.innerHTML = "STRONG";
            status.style.padding = "5px 15px";
            status.style.background = "linear-gradient(to left, rgb(0, 88, 7) 0%, #7bff00 100%)";
            status.style.color = "white";
        }
    } else {
        status.style.display = "none";
        label.style.display = "none";
    }
}