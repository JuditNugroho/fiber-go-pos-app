function formatQty(value) {
    return value.toLocaleString('en-US')
}

function qtyFormatter(value, row, index) {
    let qtyFmt = formatQty(value)
    let text_field_id = "qty_field" + index.toString()

    return `
        <div class="input-group" >
          <span class="input-group-btn">
              <button type="button" class="btn btn-default btn-qty-number" 
                      data-type="minus" data-field="${text_field_id}">
                  <span class="glyphicon glyphicon-minus"></span>
              </button>
          </span>
          <input type="text" name="${text_field_id}" data-index="${index}"
                 class="form-control input-qty-number" 
                 style="text-align: center;" 
                 min="1" max="10000" value="${qtyFmt}">
          <span class="input-group-btn">
              <button type="button" class="btn btn-default btn-qty-number" 
                      data-type="plus" data-field="${text_field_id}">
                  <span class="glyphicon glyphicon-plus"></span>
              </button>
          </span>
      </div>
    `
}

$(document).on('click', "button.btn-qty-number", function (e) {
    e.preventDefault();

    // initialization variable
    let fieldName = $(this).attr('data-field');
    let type = $(this).attr('data-type');
    let input = $("input[name='" + fieldName + "']");
    let currentVal = parseInt(input.val());

    if (!isNaN(currentVal)) {
        if (type === 'minus') {
            if (currentVal > input.attr('min')) {
                input.val(currentVal - 1).trigger("change");
            }
            if (parseInt(input.val()) === input.attr('min')) {
                $(this).attr('disabled', true);
            }

        } else if (type === 'plus') {
            if (currentVal < input.attr('max')) {
                input.val(currentVal + 1).trigger("change");
            }
            if (parseInt(input.val()) === input.attr('max')) {
                $(this).attr('disabled', true);
            }
        }
    } else {
        input.val(0);
    }
});

$(document).on('focus', "input.input-qty-number", function () {
    $(this).data('oldValue', $(this).val());
});

$(document).on('keydown', "input.input-qty-number", function (e) {
    // Allow: backspace, delete, tab, escape, enter and .
    if ($.inArray(e.keyCode, [46, 8, 9, 27, 13]) !== -1 ||
        // Allow: Ctrl+A
        (e.keyCode === 65 && e.ctrlKey === true) ||
        // Allow: home, end, left, right
        (e.keyCode >= 35 && e.keyCode <= 39)) {
        // let it happen, don't do anything
        return;
    }
    // Ensure that it is a number and stop the keypress
    if ((e.shiftKey || (e.keyCode < 48 || e.keyCode > 57)) && (e.keyCode < 96 || e.keyCode > 105)) {
        e.preventDefault();
    }
});