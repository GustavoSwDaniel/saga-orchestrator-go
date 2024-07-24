from typing import List
from pydantic import BaseModel
from datetime import datetime
from typing import List

class OrderRegister(BaseModel):
    product_uuid: str
    quantity: int
