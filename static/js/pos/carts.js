let dataCarts = []

function getAllDataCarts() {
    return dataCarts
}

function getDataCartByIndex(index) {
    return dataCarts[index]
}

function getTotalDataCarts() {
    return dataCarts.length
}

function checkDataCartFromProductID(product_id) {
    const checkProductID = element => element.product_id === product_id;
    return dataCarts.some(checkProductID)
}

function insertDataCart(data) {
    dataCarts[dataCarts.length] = data;
}

function updateQtyDataCart(index, qty) {
    dataCarts[index]["qty"] = qty;

    let data = dataCarts[index];
    dataCarts[index]["total"] = data["price_used"] * data["qty"];
}

function removeAllDataCart() {
    dataCarts = []
}
