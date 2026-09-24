import urllib.request
import json
import urllib.error
import io
import uuid

BASE_URL = 'http://localhost:8080'

# Login to get fresh token
req = urllib.request.Request(
    BASE_URL + '/auth/login',
    data=json.dumps({'no_telp': '081234567890', 'kata_sandi': 'password123'}).encode('utf-8'),
    headers={'Content-Type': 'application/json'},
    method='POST'
)
res = json.loads(urllib.request.urlopen(req).read().decode('utf-8'))
token = res['data']['token']
headers = {'token': token}

def api(method, path, data=None, content_type='application/json'):
    h = dict(headers)
    body = None
    if data is not None:
        if content_type == 'application/json':
            body = json.dumps(data).encode('utf-8')
            h['Content-Type'] = 'application/json'
        else:
            body = data
            h['Content-Type'] = content_type
    r = urllib.request.Request(BASE_URL + path, data=body, headers=h, method=method)
    try:
        with urllib.request.urlopen(r) as resp:
            return resp.status, json.loads(resp.read().decode('utf-8'))
    except urllib.error.HTTPError as e:
        return e.code, json.loads(e.read().decode('utf-8'))

print('=== 9. Create Category (As Admin now) ===')
status, res = api('POST', '/category', {'nama_category': 'Elektronik'})
print(status, res)

print('=== 10. Get All Categories ===')
status, res = api('GET', '/category')
print(status, res['data'])
cat_id = res['data'][0]['id']

print('=== 11. Create Alamat ===')
alamat_payload = {
    'judul_alamat': 'Rumah Utama',
    'nama_penerima': 'Fajar Penerima',
    'no_telp': '081234567890',
    'detail_alamat': 'Jl. Diponegoro No. 45, Bandung'
}
status, res = api('POST', '/user/alamat', alamat_payload)
print(status, 'New Alamat ID:', res['data'])
alamat_id = res['data']

print('=== 12. Get Alamat List ===')
status, res = api('GET', '/user/alamat')
print(status, res['data'])

print('=== 13. Create Product (Multipart Form-Data) ===')
boundary = '----WebKitFormBoundary' + uuid.uuid4().hex
fields = {
    'nama_produk': 'Keyboard Mekanikal Evermos',
    'category_id': str(cat_id),
    'harga_reseller': '350000',
    'harga_konsumen': '450000',
    'stok': '50',
    'deskripsi': 'Keyboard mechanical RGB switch blue'
}
buf = io.BytesIO()
for k, v in fields.items():
    buf.write(f'--{boundary}\r\n'.encode('utf-8'))
    buf.write(f'Content-Disposition: form-data; name="{k}"\r\n\r\n'.encode('utf-8'))
    buf.write(f'{v}\r\n'.encode('utf-8'))

buf.write(f'--{boundary}\r\n'.encode('utf-8'))
buf.write(b'Content-Disposition: form-data; name="photos"; filename="keyboard.png"\r\n')
buf.write(b'Content-Type: image/png\r\n\r\n')
buf.write(b'FakePNGImageContentBytes\r\n')
buf.write(f'--{boundary}--\r\n'.encode('utf-8'))

status, res = api('POST', '/product', buf.getvalue(), f'multipart/form-data; boundary={boundary}')
print(status, 'New Product ID:', res['data'])
product_id = res['data']

print('=== 14. Get Product By ID ===')
status, res = api('GET', f'/product/{product_id}')
print(status, res['data']['nama_produk'], 'Stok:', res['data']['stok'], 'Photos:', len(res['data'].get('photos', [])))

print('=== 15. Create Trx (Buy 2 units of the product) ===')
trx_payload = {
    'method_bayar': 'bca',
    'alamat_kirim': alamat_id,
    'detail_trx': [
        {
            'product_id': product_id,
            'kuantitas': 2
        }
    ]
}
status, res = api('POST', '/trx', trx_payload)
print(status, 'New Trx ID:', res['data'])
trx_id = res['data']

print('=== 16. Verify Stock Deduction ===')
status, res = api('GET', f'/product/{product_id}')
print('Stock after buying 2 units:', res['data']['stok'], '(expected 48)')

print('=== 17. Get Trx By ID (Verifying LogProduk & Invoice) ===')
status, res = api('GET', f'/trx/{trx_id}')
print(status, 'Invoice:', res['data']['kode_invoice'], 'Total:', res['data']['harga_total'])
print('Detail Trx Product:', res['data']['detail_trx'][0]['product']['nama_produk'])

print('\nALL TESTS COMPLETED SUCCESSFULLY!')
