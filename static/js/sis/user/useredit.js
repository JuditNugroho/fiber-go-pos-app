function editUser(row) {
    fillFormEdit(row)
    $('#modalUpsert').modal('show');
}

function fillFormEdit(row){
    let isAdminFmt = row["is_admin"] === true ? "Ya" : "Tidak"
    $("#modalUpsert #user_id").val(row["user_id"]);
    $("#modalUpsert #user_name").val(row["user_name"]);
    $("#modalUpsert #full_name").val(row["full_name"]);
    $("#modalUpsert #password").val(row["password"]);
    $("#modalUpsert #is_admin").val(isAdminFmt);
}