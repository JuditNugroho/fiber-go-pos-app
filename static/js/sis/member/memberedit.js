function editMember(row) {
    fillFormEdit(row)
    $('#modalUpsert').modal('show');
}

function fillFormEdit(row){
    $("#modalUpsert #id").val(row["id"]);
    $("#modalUpsert #name").val(row["name"]);
    $("#modalUpsert #phone").val(row["phone"]);
}