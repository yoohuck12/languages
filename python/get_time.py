import time
from datetime import datetime

from pytz import timezone


def get_kst_time():
    print(datetime.now(timezone("Asia/Seoul")))


def time_diff():
    start = time.time()
    end = time.time()
    elapsed = end-start
    print(time.strftime("%H-%M-%S", time.gmtime(elapsed)))
