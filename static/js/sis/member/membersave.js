$("#upsertMember").on("click", function (event) {
    event.preventDefault();

    let param = getParamValue();
    if (param.err !== null) {
        alertify.alert('Pesan Peringatan', param.err);
        return
    }

    saveMember(param.data);
});

function getParamValue() {
    let id = $("#modalUpsert #id").val().trim()
    let name = $("#modalUpsert #name").val().trim()
    let phone = $("#modalUpsert #phone").val().trim()

    if (id === "") {
        return {"data": null, "err": "ID member tidak boleh kosong !"}
    }

    if (name === "") {
        return {"data": null, "err": "Nama member tidak boleh kosong !"}
    }

    return {
        "err": null,
        "data": {
            "id": id,
            "name": name,
            "phone": phone,
        }
    }

}

function saveMember(data) {
    let baseURL = $('#baseURL').text();
    $.ajax({
        type: "POST",
        async: false,
        data: JSON.stringify(data),
        contentType: 'application/json',
        url: baseURL + "svc/member/upsert",
    }).then(function (res) {
        clearFormInput();
        $("#modalUpsert").modal('toggle');
        alertify.success("Data berhasil disimpan");
        $('#table').bootstrapTable('refresh');
    }).catch(function (a) {
        alertify.alert('Pesan Error', a.responseText);
    });
}