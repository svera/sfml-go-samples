package animatedSprite

import (
    sf "bitbucket.org/krepa098/gosfml2"
    "sergio/animated_sprites/animation"
    "math"
)

type AnimatedSprite struct {
    m_animation *animation.Animation
    m_frameTime uint
    m_currentTime uint
    m_currentFrame uint8
    m_isPaused bool
    m_isLooped bool
    m_texture *sf.Texture
    m_vertices [4]sf.Vertex
}

func NewAnimatedSrite(frameTime uint, paused bool, looped bool) AnimatedSprite {
    a := new(AnimatedSprite)
    a.m_animation = nil
    a.m_frameTime = frameTime
    a.m_currentFrame = 0
    a.m_isPaused = paused
    a.m_isLooped = looped
    return *a
}

func (this *AnimatedSprite) Update(deltaTime uint) {
// if not paused and we have a valid animation
    if !this.m_isPaused && this.m_animation != nil {
        // add delta time
        this.m_currentTime += deltaTime

        // if current time is bigger then the frame time advance one frame
        if this.m_currentTime >= this.m_frameTime {
            // reset time, but keep the remainder
            this.m_currentTime = this.m_currentTime % this.m_frameTime

            // get next Frame index
            if this.m_currentFrame + 1 < this.m_animation.GetSize() {
                this.m_currentFrame++
            } else {
                // animation has ended
                this.m_currentFrame = 0 // reset to start

                if !this.m_isLooped {
                    this.m_isPaused = true
                }

            }

            // set the current frame, not reseting the time
            this.SetFrame(this.m_currentFrame, false)
        }
    }
}

func (this *AnimatedSprite) SetAnimation(animation *animation.Animation) {
    this.m_animation = animation
    this.m_texture = this.m_animation.GetSpriteSheet()
    this.m_currentFrame = 0
    this.SetFrame(this.m_currentFrame, true)
}

func (this *AnimatedSprite) SetFrameTime(time uint) {
    this.m_frameTime = time
}

func (this *AnimatedSprite) Play(animation ...*animation.Animation) {
    if len(animation) > 0 && this.GetAnimation() != animation[0] {
        this.SetAnimation(animation[0])
    }
    this.m_isPaused = false
}

func (this *AnimatedSprite) Pause() {
    this.m_isPaused = true
}

func (this *AnimatedSprite) Stop() {
    this.m_isPaused = true
    this.m_currentFrame = 0;
    this.SetFrame(this.m_currentFrame, true)
}

func (this *AnimatedSprite) SetLooped(looped bool) {
    this.m_isLooped = looped
}

func (this *AnimatedSprite) SetColor(color sf.Color) {
    this.m_vertices[0].Color = color
    this.m_vertices[1].Color = color
    this.m_vertices[2].Color = color
    this.m_vertices[3].Color = color
}

func (this *AnimatedSprite) GetAnimation() *animation.Animation {
    return this.m_animation
}

func (this *AnimatedSprite) GetLocalBounds() sf.FloatRect {
    rect := this.m_animation.GetFrame(this.m_currentFrame)
    width := float32(math.Abs(float64(rect.Width)))
    height := float32(math.Abs(float64(rect.Height)))

    return sf.FloatRect{0.0, 0.0, width, height}
}

func (this *AnimatedSprite) GetGlobalBounds() sf.FloatRect {
    transform := sf.NewTransformable().GetTransform()
    return transform.TransformRect(this.GetLocalBounds())
}

func (this *AnimatedSprite) IsLooped() bool {
    return this.m_isLooped
}

func (this *AnimatedSprite) IsPlaying() bool {
    return !this.m_isPaused
}

func (this *AnimatedSprite) GetFrameTime() uint {
    return this.m_frameTime
}

func (this *AnimatedSprite) SetFrame(newFrame uint8, resetTime bool) {
    if this.m_animation != nil {
        //calculate new vertex positions and texture coordiantes
        rect := this.m_animation.GetFrame(newFrame)

        this.m_vertices[0].Position = sf.Vector2f{0.0, 0.0}
        this.m_vertices[1].Position = sf.Vector2f{0.0, float32(rect.Height)}
        this.m_vertices[2].Position = sf.Vector2f{float32(rect.Width), float32(rect.Height)}
        this.m_vertices[3].Position = sf.Vector2f{float32(rect.Width), 0.0}

        left := float32(rect.Left) + 0.0001
        right := left + float32(rect.Width)
        top := float32(rect.Top)
        bottom := top + float32(rect.Height)

        this.m_vertices[0].TexCoords = sf.Vector2f{left, top}
        this.m_vertices[1].TexCoords = sf.Vector2f{left, bottom}
        this.m_vertices[2].TexCoords = sf.Vector2f{right, bottom}
        this.m_vertices[3].TexCoords = sf.Vector2f{right, top}
    }

    if (resetTime) {
        this.m_currentTime = 0.0
    }
}

/*
class AnimatedSprite : public sf::Drawable, public sf::Transformable
{
public:
    explicit AnimatedSprite(sf::Time frameTime = sf::seconds(0.2f), bool paused = false, bool looped = true);

    void update(sf::Time deltaTime);
    void setAnimation(const Animation& animation);
    void setFrameTime(sf::Time time);
    void play();
    void play(const Animation& animation);
    void pause();
    void stop();
    void setLooped(bool looped);
    void setColor(const sf::Color& color);
    const Animation* getAnimation() const;
    sf::FloatRect getLocalBounds() const;
    sf::FloatRect getGlobalBounds() const;
    bool isLooped() const;
    bool isPlaying() const;
    sf::Time getFrameTime() const;
    void setFrame(std::size_t newFrame, bool resetTime = true);

private:
    const Animation* m_animation;
    sf::Time m_frameTime;
    sf::Time m_currentTime;
    std::size_t m_currentFrame;
    bool m_isPaused;
    bool m_isLooped;
    const sf::Texture* m_texture;
    sf::Vertex m_vertices[4];

    virtual void draw(sf::RenderTarget& target, sf::RenderStates states) const;

};
*/