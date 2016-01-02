# Introduction #

Everyone is welcome to propose code modifications and additions to Pokemon Universe. No matter if it's for the client, server or any of our packages. However, a strict procedure has to be followed to do this. This page will instruct you on how to contribute to Pokemon Universe.

# Prerequisites #

You will need the following:
  * Subversion
  * Python
  * A Google account (used for gmail, youtube etc.)

# Contributing code #
## Step 1. Obtaining the source ##
Use your Subversion client to check out the trunk of our source
```
svn checkout http://pokemon-universe.googlecode.com/svn/trunk/
```

## Step 2. Add your changes to the code ##
Simple edit the files or add new files. Be sure to "svn add" any new files.

## Step 3. Download upload.py ##
Download upload.py here:
http://codereview.appspot.com/static/upload.py
And save it in the root of the Pokemon Universe source trunk

## Step 4. Run upload.py ##
Run upload.py on your modified/added files. For example if you edited Client/main.go the command is:
```
python upload.py Client/main.go
```
When asked, supply a title for your modification and enter your Google account credentials. The output will be a link to your generated code review page. It will look like http://codereview.appspot.com/#######
Visit the page, log in and edit the issue to add a description. Also be sure to add pu.urmel@gmail.com and mr\_dark@darkweb.nl as reviewers.

## Step 5. Keep track of the issue ##
Pokemon Universe developers will at some point review your patch and possibly add comments. If there are any required changes you can apply those changes to your code and use upload.py again with the added "-i <your issue number here>" parameter. Your new patch set will be uploaded to your issue to be reviewed. This process will continue until a developer approves of your change and commits it to the svn.

**NOTE: All patches have to be reviewed, there are no exceptions. Please do not mail any developers with patches, but always use this procedure.**