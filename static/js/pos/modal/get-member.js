$(function () {
    initTblMember();
})

function addMemberToCart(row) {
    // set value to button add member
    let btnAddMember = $("#btnAddMember")
    btnAddMember.html(row["name"]);
    btnAddMember.prop('disabled', true);

    // set member text input
    $("#member-id").val(row["id"])
    $("#member-name").val(row["name"])

    $('#modalAddMember').modal('hide');
}

function addMemberFmt() {
    return '<a class="add-member" title="Tambah Member"><i class="fa fa-plus-circle"></i></a>'
}

window.addMemberEventAction = {
    'click .add-member': function (e, value, row, index) {
        let textString = 'Apakah anda yakin ingin menambahkan Member ' + row["name"] + ' ?'
        alertify.confirm('Dialog Konfirmasi', textString, function () {
            addMemberToCart(row)
        }, null).setting({'labels': {ok: 'Ya', cancel: 'Tidak'}});
    }
}

function initTblMember() {
    $('#table-add-member').bootstrapTable({
        locale: $('#locale').val(),
        formatNoMatches: function () {
            return 'No data found';
        },
        columns: [
            [
                {
                    width: 150,
                    field: 'id',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'ID Member',
                },
                {
                    width: 350,
                    field: 'name',
                    align: 'left',
                    widthUnit: "px",
                    valign: 'middle',
                    title: 'Nama Member',
                },
                {
                    width: 150,
                    title: 'Action',
                    field: 'action',
                    align: 'center',
                    clickToSelect: false,
                    formatter: addMemberFmt,
                    events: window.addMemberEventAction,
                }
            ],
        ]
    });
}

function headerTblAddMember(column) {
    return {
        id: {
            css: {
                background: '#34b9f9',
                color: '#FFFFFF'
            }
        },
        name: {
            css: {
                background: '#34b9f9',
                color: '#FFFFFF'
            }
        },
        action: {
            css: {
                background: '#34b9f9',
                color: '#FFFFFF'
            }
        }
    }[column.field]
}

// your custom ajax request here
function ajaxAddMemberRequest(params) {
    let baseURL = $('#baseURL').text();
    $.ajax({
        'method': "GET",
        'url': baseURL + "svc/dt_members",
        'contentType': 'application/json',
    }).done(function (data) {
        $('#table-add-member').bootstrapTable('resetView');
        params.success(data);
    });
}
