fb_photo_downloader
======================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/toomore/gogrs/master/LICENSE)

A facebook page photo album tool that supports concurrency download. This tool help you to download those photos for your backup, all the photos still own by original creator.

Install
--------------

    go get -u -x github.com/PavelPavells/fb_photo_downloader

Note, you need go to [faceook developer page](https://developers.facebook.com/tools/explorer?method=GET&path=me) to get latest token and update in environment variables.

     export FBTOKEN = "YOUR_TOKEN_HERE"

Usage
---------------------

    fb_photo_downloader [options] 

All the photos will download to `USERS/Pictures/fb_photo_downloader` and it will separate folder by page name and album name.

Options
---------------

- `-n` Facebook page name such as: [scottiepippen](https://www.facebook.com/scottiepippen), or input facebook id if you know such as 112743018776863 
- `-c` number of workers. (concurrency), default workers is "2"


Examples
---------------

Download all photos from Scottie Pippen facebook pages with 10 workers.

  fb_photo_downloader -n=scottiepippen -c=10


TODOs
---------------

- Support specific album download.
- Support firend/self album download for backup.
- Support CLI to select which album you want to download.


Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`


Advertising
---------------

If you want to browse facebook page on your iPhone, why not check my App here :p [粉絲相簿](https://itunes.apple.com/tw/app/fen-si-xiang-bu/id839324997?l=zh&mt=8)


License
---------------

This package is licensed under MIT license. See LICENSE for details.
