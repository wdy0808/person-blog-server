#!/bin/bash
function echo_stderr {
    printf '%s\n' "${1}" #> /dev/stderr
}

function __usage {
	echo_stderr "usage: pull code to latest, run or stop project code"
	echo_stderr "-project <project name>"
	echo_stderr "-branch <branch name, default master>"
	echo_stderr "-cmd <start, stop, restart>"
	echo_stderr "-help, print this message"
}

declare -a MODULE_VERSION
BUILD_PROJECT=
BRANCH="master"
CMD=

while [[ ${1} ]]
do
    case ${1} in 
        -project)
            shift
            BUILD_PROJECT=${1}
            ;;
        -branch)
            shift
            BRANCH=${1}
            ;;
        -cmd)
            shift
            CMD=${1}
            ;;
        -help)
            shift
            __usage
            exit 0
            ;;
        *)
            MODULE_VERSION=(${MODULE_VERSION[@]} ${1})
            ;;
    esac
    shift
done

if [ -z "${BUILD_PROJECT}" ]
then
    echo_stderr "illegal build project"
    __usage
    exit 1
fi

cd ${PROJECT_PATH}/${BUILD_PROJECT}

CURRENT_BRANCH=`git symbolic-ref --short -q HEAD`
if [ ${BRANCH} != ${CURRENT_BRANCH} ]
then
	if_branch_exist=`git branch | grep -w ${BRANCH}`
	if [ -n "${if_branch_exist}" ]
	then
		git checkout ${BRANCH}
	else
		git checkout -b ${BRANCH}
	fi
fi

PULL_RESULT=`git pull origin ${BRANCH} | grep "fatal: Couldn't find remote ref"`
if [ -n "${PULL_RESULT}" ]
then
	echo_stderr "back to branch ${CURRENT_BRANCH}"
	git checkout ${CURRENT_BRANCH}
	exit 1
fi

if [ -z "${CMD}" ]
then
	exit 0
fi

function mustMkdir {
	if [ ! -d "${1}" ]
	then
		mkdir ${1}
	fi
}

function storeDependencies {
	mustMkdir ${RUNPATH}/${BUILD_PROJECT}/version
	declare -a depencdencies
	go list -m all | head -n 1 > ${RUNPATH}/${BUILD_PROJECT}/version/module
	go list -m all | sed '1d' > ${RUNPATH}/${BUILD_PROJECT}/version/build_dependencies
}

function start {
	go install
	mustMkdir ${RUNPATH}/${BUILD_PROJECT}
	mustMkdir ${RUNPATH}/${BUILD_PROJECT}/log
	if [ -f "${RUNPATH}/${BUILD_PROJECT}/log/logfile" ]
	then 
		mv ${RUNPATH}/${BUILD_PROJECT}/log/logfile ${RUNPATH}/${BUILD_PROJECT}/log/logfile.backup
	fi
	mv -f ${GOBIN}/${BUILD_PROJECT} ${RUNPATH}/${BUILD_PROJECT}/${BUILD_PROJECT}
	cp -R -f ${PROJECT_PATH}/${BUILD_PROJECT}/conf ${RUNPATH}/${BUILD_PROJECT}/conf
	${RUNPATH}/${BUILD_PROJECT}/${BUILD_PROJECT} & 2>${RUNPATH}/${BUILD_PROJECT}/log/logfile
	storeDependencies
}

function stop {
	pid=`ps -ef | grep "${RUNPATH}/${BUILD_PROJECT}/${BUILD_PROJECT}" | grep -v grep | awk '{print $2}'`
	kill -n 15 ${pid}
}

function restart {
	stop
	start
}

case ${CMD} in
	"start")
		start
		;;
	"stop")
		stop
		;;
	"restart")
		restart
		;;
	*)
		echo_stderr "unknown cmd [${CMD}]"
		exit 1
		;;
esac
