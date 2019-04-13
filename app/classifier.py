import torch
import torchvision.transforms as transforms
import trf
from vgg import VGG19
from PIL import Image


def classifier(img):
    img = img.convert('RGB')
    img = trf.resize(img, (64, 64))
    img = transforms.Compose([
        transforms.ToTensor(),
        transforms.Normalize(mean=(0.5071, 0.4867, 0.4408),
                             std=(0.2675, 0.2565, 0.2761))
    ])(img)
    img = torch.reshape(img, (1, 3, 64, 64))
    net = VGG19(num_class=172)
    net.load_state_dict(torch.load("ckpt.pth"))
    net.eval()
    with torch.no_grad():
        outputs = net(img)
        pred = outputs.max(1)[1]
    return int(pred)
