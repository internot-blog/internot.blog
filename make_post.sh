#!/usr/bin/env bash

# pushd ./internot.blog

# remove old posts
rm -rf ./posts/*

# generate new post
go run main.go

# popd
pushd ../

# copy posts over to site repo
cp -r ./internot.blog/posts/* ./internot-blog.github.io/content/post

pushd ./internot-blog.github.io/content/post

git add .

GIT_COMMITTER_NAME="internot.blog" GIT_COMMITTER_EMAIL="bot@internot.blog" git commit -m "Added new post (automatic)" --author="internot.blog <bot@internot.blog>"
git push

popd
popd
