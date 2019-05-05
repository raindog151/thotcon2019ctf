#!/usr/bin/env python
# Socket client example in python

import socket
import sys
import math

host = '3.14.124.98'
port = 1337

# create socket
print('# Creating socket')
#try:
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
#except socket.error:
#    print('Failed to create socket')
#    sys.exit()

# print('# Getting remote IP address')
#try:
#    remote_ip = socket.gethostbyname( host )
#except socket.gaierror:
#    print('Hostname could not be resolved. Exiting')
#    sys.exit()

# Connect to remote server
print('# Connecting to server, ' + host)
s.connect((host, port))


while True:

  # Receive data
  print('# Receive data from server')
  reply = s.recv(256)

  print reply

  # Send data to remote server
  # print('# Sending data to server')
  # request =

  #print str(eval(reply))

  #try:

  if reply.strip() == 'sqrt(4)':
    s.sendall('2\n')
  elif reply.strip() == 'x=1; y=2; x+y':
    s.sendall('3\n')
  elif reply.strip() == 'a=666; b=2; a*=2; b+=3; a+b':
    s.sendall('1337\n')
  else:
    s.sendall(str(eval(reply)) + '\n')
  #except socket.error:
  #    print 'Send failed'
  #    sys.exit()

  continue
