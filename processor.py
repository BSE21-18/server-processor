import pickle

from flask import Flask,request
import json

app = Flask(__name__)

@app.route('/', methods=['GET'])
def home():
    return '''<h1>DATAVOC SENSOR ARRAY PROCESSOR </h1> <p> Welcome to DATAVOC Machine Learning API </p>'''
  
@app.route('/processor', methods=['POST'])
def predict():
    event = json.loads(request.data)
    print(event)
    device = event["sniffer"]
    values = event["values"]
    Ir = None
    with open('model_pkl','rb') as f:
        Ir = pickle.load(f)
    res = Ir.predict(values)
    print(res)
    response = '{"sniffer": "'+device+'", "value": '+res+'}'
    return str(response)
    
if __name__ == '__main__':
    app.run(debug=True)
    
    
    
