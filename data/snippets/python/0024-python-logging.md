## Logging
Structured logging with logging module
```python
import logging

# Basic logging setup
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s',
    handlers=[
        logging.FileHandler('app.log'),
        logging.StreamHandler()
    ]
)

logger = logging.getLogger(__name__)

def demo_logging():
    logger.debug("Debug message (won't show)")
    logger.info("Application started")
    logger.warning("This is a warning")
    logger.error("This is an error")
    
    try:
        result = 10 / 0
    except ZeroDivisionError:
        logger.exception("Division by zero occurred")

# Custom logger with different level
file_logger = logging.getLogger('file_logger')
file_logger.setLevel(logging.DEBUG)

file_handler = logging.FileHandler('debug.log')
file_handler.setFormatter(
    logging.Formatter('%(name)s - %(levelname)s - %(message)s')
)
file_logger.addHandler(file_handler)

# Usage
demo_logging()
file_logger.debug("This goes to debug.log")
logger.info("App running successfully")
```