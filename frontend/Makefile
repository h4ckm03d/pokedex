deployfe:
	rm -rf dist && \
	yarn build && \
	aws s3 cp dist s3://pokedextcg --recursive --grants read=uri=http://acs.amazonaws.com/groups/global/AllUsers && \
	aws s3 website s3://pokedextcg --index-document index.html
