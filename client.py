import socket,time

def main():
    # 设置服务器的IP地址和端口号
    server_ip = "localhost"
    server_port = 8080

    # 创建TCP套接字
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    try:
        # 连接服务器
        client_socket.connect((server_ip, server_port))
        print("Connected to server at {}:{}".format(server_ip, server_port))

        # 发送数据给服务器
        message = "Hello, server!"
        client_socket.sendall(message.encode())

        # 接收服务器的响应
        response = client_socket.recv(1024).decode()
        print("Response from server:", response)

    except ConnectionRefusedError:
        print("Error: Connection refused. Make sure the server is running.")

    finally:
        time.sleep(10)
        # 关闭套接字连接
        client_socket.close()

if __name__ == "__main__":
    main()
