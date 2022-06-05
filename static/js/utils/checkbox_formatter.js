function checkboxFormatter(data) {
    let statusCheckbox = data ? 'checked' : '';
    return '<input style="vertical-align: center;horiz-align: center;" type="checkbox" onclick="return false" ' + statusCheckbox + '>'
}