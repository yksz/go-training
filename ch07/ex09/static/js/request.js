function requestSort(key) {
  var form = document.createElement('form');
  form.action = '/'
  form.method = 'GET'

  var param = document.createElement('input');
  param.name = 'sort';
  param.value = key;
  form.appendChild(param);

  form.submit()
}
