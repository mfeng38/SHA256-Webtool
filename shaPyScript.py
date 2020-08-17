import subprocess, sys

arg = sys.argv[1].encode('utf-8').hex()
s = 'export GOPATH=$HOME/project; go build main.go; ./main ' + arg
p = subprocess.Popen(s, cwd=r'src/main', shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
output, err = p.communicate()
if err.decode('ascii') == '':
    print(output.decode('ascii'))
else:
    print(err.decode('ascii'))
