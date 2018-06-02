echo "your synbolic link"

PDF_ROOT=$(cd $(dirname $0)/; pwd)

FROM=${PDF_ROOT}/wkhtmltox/bin/wkhtmltopdf
echo "ln -s ${FROM} /usr/local/bin"
