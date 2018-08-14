#!/bin/sh

C=go-package-plantuml
GOPATH=/gp
PROJECT=/gp/src/github.com/hyperledger/fabric-sdk-go/
OUTPUTDIR=~/docs/plantuml
plantjmlJar=/gp/bin/plantuml.jar

NODENAME=$1

for showtest in true false
do
    for x in 1
    do
        echo "$C --codedir $PROJECT --gopath $GOPATH --outputdir $OUTPUTDIR \
             --nodename NODENAME --nodedepth $x \
             --ignoredir $PROJECT/internal \
             --ignoredir $PROJECT/third_party \
            --ignorenode closable \
            --ignorenode closeable \
            --testpartialdir test \
            --testpartialdir mocks \
            --showtest $showtest"

        $C --codedir $PROJECT --gopath $GOPATH --outputdir $OUTPUTDIR \
            --nodename $NODENAME --nodedepth $x \
            --ignoredir $PROJECT/internal \
            --ignoredir $PROJECT/third_party \
            --ignorenode closable \
            --ignorenode closeable \
            --testpartialdir test \
            --testpartialdir mocks \
            --showtest $showtest
        sed -i 's/github.com\\\\hyperledger\\\\//g' $OUTPUTDIR/node-$NODENAME-$x-$showtest.puml
        java -Xmx2048m -jar /gp/bin/plantuml.jar $OUTPUTDIR/node-$NODENAME-$x-$showtest.puml -tsvg
    done
done
