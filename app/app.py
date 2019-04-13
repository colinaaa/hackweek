from flask import Flask, jsonify, request
from recommend import recommend
import requests
from classifier import classifier
from PIL import Image
app = Flask(__name__)


@app.route("/api/photo", methods=["POST"])
def photo():
    try:
        img = request.get_data()
    except Exception as e:
        print(e)
        return jsonify({"msg": "photo: get_data wrong", "statue_code": 500})
    pic = Image.frombytes('RGBA', (128, 128), img, 'raw')
    # with open("photo.jpg", "wb") as photo:
    #     try:
    #         photo.write(request.get_data())
    #     except Exception as e:
    #         print(e)
    #         photo.close()
    #         return jsonify({"msg": "server wrong", "statue_code": 500})
    #     finally:
    #         photo.close()
    res = classifier(pic)
    print(res)
    with open("food_list.txt") as txt:
        food = txt.readlines()[res]
    return jsonify({"res": food, "statue_code": 200})


@app.route("/api/user/<num>/recommend")
def rec(num):
    try:
        headers = request.headers
    except:
        print("no token")
        return jsonify({"msg": "broken token", "statue_code": 403})

    try:
        resp = requests.get(
            'http://localhost/api/user/{}/diet'.format(str(num)),
            headers=headers)
    except:
        print("request wrong")
        return jsonify({"msg": "server wrong", "statue_code": 500})
    try:
        foods = resp.json()
    except Exception as e:
        print(e)
        return jsonify({"msg": "server wrong", "statue_code": 500})
    print(type(foods))
    print(foods)
    try:
        res = recommend(foods)
    except Exception as e:
        raise e
    return jsonify(res)
