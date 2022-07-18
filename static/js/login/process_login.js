function processLogin() {
    let baseURL = $('#baseURL').text();

    let user_name = $("#user_name").val().trim();
    if (user_name === "") {
        alertify.alert("Pesan Dialog", "Harap masukkan username");
        return false
    }

    let password = $("#password").val().trim();
    if (password === "") {
        alertify.alert("Pesan Dialog", "Harap masukkan password");
        return false
    }

    $.ajax({
        type: "POST",
        async: false,
        url: baseURL + "svc/login",
        contentType: 'application/json',
        data: JSON.stringify({"user_name": user_name, "password": password}),
    }).then(function (res) {
        sessionStorage.setItem("user_id", res["user_id"]);
        sessionStorage.setItem("user_name", res["user_name"]);
        window.location.href = baseURL + "pos";
    }).catch(function (a) {
        let res = a.responseJSON;
        alertify.alert('Pesan Error', res["message"]);
    });
}