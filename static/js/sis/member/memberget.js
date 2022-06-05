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
        editMember(row);
    }, 'click .remove': function (e, value, row, index) {
        alertify.confirm('Dialog Konfirmasi', 'Apakah anda yakin ingin menghapus data ini?', function () {
            deleteMember(row);
        }, null).setting({'labels': {ok: 'Ya', cancel: 'Tidak'}});
    }
}

function initTable() {
    $('#table').bootstrapTable({
        locale: $('#locale').val(),
        columns: [
            [
                {
                    width: 150,
                    field: 'id',
                    title: 'ID',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 350,
                    title: 'Nama',
                    field: 'name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                },
                {
                    width: 250,
                    align: 'left',
                    title: 'No HP',
                    field: 'phone',
                    widthUnit: "px",
                    valign: 'middle',
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
        'url': baseURL + "svc/dt_members",
        'contentType': 'application/json',
    }).done(function (data) {
        $('#table').bootstrapTable('resetView');
        params.success(data);
    });
}

function clearFormInput() {
    $("#modalUpsert #id").val("");
    $("#modalUpsert #name").val("");
    $("#modalUpsert #phone").val("");
}