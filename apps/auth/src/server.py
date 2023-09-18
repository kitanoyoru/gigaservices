import grpc

from futures import ThreadPoolExecutor



class Server:
	def __init__(self, config: Config):
		self._config = config

	def serve(self):
		server = grpc.server(ThreadPoolExecutor(max_workers=config.max_grpc_workers))
k
