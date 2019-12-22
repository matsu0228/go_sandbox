
show:


deploy:
	gcloud config set project reactfiredemo-96257
	# # to upper 
	# gcloud functions deploy MakeUpperCase \
	# 	--runtime go111 \
	# 	--trigger-event providers/cloud.firestore/eventTypes/document.write \
	# 	--trigger-resource "projects/reactfiredemo-96257/databases/(default)/documents/messages/{pushId}"
	# trans
	gcloud functions deploy Translate \
		--runtime go111 \
		--trigger-event providers/cloud.firestore/eventTypes/document.write \
		--trigger-resource "projects/reactfiredemo-96257/databases/(default)/documents/translate/{pushId}"
