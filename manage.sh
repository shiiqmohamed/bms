#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

case "$1" in
  start)
    echo -e "${YELLOW}📦 Loading environment variables...${NC}"
    source setenv.sh
    
    # Check if already running
    if pgrep -f bms_api > /dev/null; then
        echo -e "${RED}❌ BMS API is already running${NC}"
        exit 1
    fi
    
    echo -e "${YELLOW}🚀 Starting BMS API on port $SERVER_PORT...${NC}"
    nohup ./bms_api > bms_output.log 2>&1 &
    sleep 3
    
    if pgrep -f bms_api > /dev/null; then
        PID=$(pgrep -f bms_api)
        echo -e "${GREEN}✅ BMS API started successfully (PID: $PID)${NC}"
        echo -e "${GREEN}📡 API URL: http://localhost:$SERVER_PORT${NC}"
        echo -e "${GREEN}📝 Logs: tail -f bms_output.log${NC}"
        
        # Test database connection
        sleep 2
        if grep -q "Database connected successfully" bms_output.log; then
            echo -e "${GREEN}✅ Database connection successful${NC}"
        else
            echo -e "${RED}❌ Database connection failed - check logs${NC}"
            tail -n 10 bms_output.log
        fi
    else
        echo -e "${RED}❌ Failed to start. Check logs:${NC}"
        tail -n 30 bms_output.log
    fi
    ;;
    
  stop)
    echo -e "${YELLOW}🛑 Stopping BMS API...${NC}"
    pkill -f bms_api
    sleep 2
    if pgrep -f bms_api > /dev/null; then
        echo -e "${RED}❌ Failed to stop${NC}"
    else
        echo -e "${GREEN}✅ BMS API stopped${NC}"
    fi
    ;;
    
  restart)
    $0 stop
    sleep 3
    $0 start
    ;;
    
  status)
    if pgrep -f bms_api > /dev/null; then
        PID=$(pgrep -f bms_api)
        echo -e "${GREEN}✅ BMS API is running (PID: $PID)${NC}"
        
        # Check if port is listening
        PORT=$(netstat -tlnp 2>/dev/null | grep $PID | awk '{print $4}' | cut -d: -f2 | head -1)
        if [ -n "$PORT" ]; then
            echo -e "${GREEN}📡 Listening on port: $PORT${NC}"
        else
            echo -e "${RED}❌ Not listening on any port${NC}"
        fi
        
        # Show last 5 lines of log
        echo -e "${YELLOW}📝 Last 5 log entries:${NC}"
        tail -5 bms_output.log
    else
        echo -e "${RED}❌ BMS API is not running${NC}"
    fi
    ;;
    
  logs)
    tail -f bms_output.log
    ;;
    
  test)
    echo -e "${YELLOW}Testing local API...${NC}"
    curl -s http://localhost:$SERVER_PORT/health | python3 -m json.tool 2>/dev/null || echo -e "${RED}❌ API not responding${NC}"
    ;;
    
  test-remote)
    echo -e "${YELLOW}Testing remote API (bms.somaalict.com)...${NC}"
    curl -s http://bms.somaalict.com/health | python3 -m json.tool 2>/dev/null || echo -e "${RED}❌ Remote API not responding${NC}"
    ;;
    
  debug)
    echo -e "${YELLOW}Running in debug mode (Ctrl+C to stop)...${NC}"
    source setenv.sh
    ./bms_api
    ;;
    
  *)
    echo "Usage: $0 {start|stop|restart|status|logs|test|test-remote|debug}"
    ;;
esac
