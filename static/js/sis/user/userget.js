$(function () {
    initTable();
})

function actionFormatter() {
    return [
        '<a class="edit" href="javascript:void(0)" title="Edit"><i class="fa fa-edit"></i></a>',
        '<a class="remove" href="javascript:void(0)" title="Remove"><i class="fa fa-trash"></i></a>'
    ].join('')
}

window.eventActions = {
    'click .edit': function (e, value, row, index) {
        editUser(row);
    }, 'click .remove': function (e, value, row, index) {
        alertify.confirm('Dialog Konfirmasi', 'Apakah anda yakin ingin menghapus data ini?', function () {
            deleteUser(row);
        }, null).setting({'labels': {ok: 'Ya', cancel: 'Tidak'}});
    }
}

function isAdminFormatter(data) {
    let statusCheckbox = data ? 'checked' : '';
    return '<input style="vertical-align: center;horiz-align: center;" type="checkbox" onclick="return false" ' + statusCheckbox + '>'
}

function initTable() {
    $('#table').bootstrapTable({
        locale: $('#locale').val(),
        columns: [
            [
                {
                    width: 150,
                    align: 'left',
                    widthUnit: "px",
                    field: 'user_id',
                    title: 'User ID',
                    valign: 'middle',
                },
                {
                    width: 200,
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Username',
                    field: 'user_name',
                },
                {
                    width: 250,
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Full Name',
                    field: 'full_name',
                },
                {
                    width: 100,
                    title: 'User Admin',
                    field: 'is_admin',
                    widthUnit: "px",
                    align: 'center',
                    valign: 'middle',
                    formatter: isAdminFormatter
                },
                {
                    title: 'Action',
                    align: 'center',
                    clickToSelect: false,
                    formatter: actionFormatter,
                    events: window.eventActions,
                }
            ],
        ]
    });
}

function ajaxRequest(params) {
    let baseURL = $('#baseURL').text();
    $.ajax({
        'method': "GET",
        'url': baseURL + "svc/dt_users",
        'contentType': 'application/json',
    }).done(function (data) {
        $('#table').bootstrapTable('resetView');
        params.success(data);
    });
}

function clearFormInput() {
    $("#modalUpsert #user_id").val("");
    $("#modalUpsert #user_name").val("");
    $("#modalUpsert #full_name").val("");
    $("#modalUpsert #password").val("");
    $("#modalUpsert #is_admin").val("Ya");
}