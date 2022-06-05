document.querySelector('#search_product').addEventListener('keypress', function (e) {
    e.stopImmediatePropagation();
    if (e.key === 'Enter') {
        getProductFromTextField(this.value)
    }
});


function getProductFromTextField(search) {
    let baseURL = $('#baseURL').text();
    $.ajax({
        'method': "GET",
        'url': baseURL + "svc/product",
        'contentType': 'application/json',
        "data": {
            "search": search,
        },
    }).done(function (data) {
        addProductFromSearch(data)
    }).catch(function (a) {
        alertify.alert('Pesan Dialog', buildErrorMessage(a.responseJSON["message"]));
    });
}

function addProductFromSearch(row) {
    $('#search_product').val('');
    addProductToTblCart(row)
}